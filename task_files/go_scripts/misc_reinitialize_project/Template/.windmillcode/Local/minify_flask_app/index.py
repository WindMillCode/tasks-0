import os
import shutil
import python_minifier
import asyncio
import aiofiles

source_folder = os.path.join("apps","backend","FlaskApp")
destination_folder = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'FlaskApp')


def should_skip(path):
  return 'site-packages' in path or 'unit_tests' in path

async def minify_file(source_path, dest_path, failed_files):
  try:
    async with aiofiles.open(source_path, 'r') as f:
      code = await f.read()
    minified_code = python_minifier.minify(code)
    os.makedirs(os.path.dirname(dest_path), exist_ok=True)
    async with aiofiles.open(dest_path, 'w') as f:
      await f.write(minified_code)
  except Exception as e:
    print(type(e).__name__)
    print(f"Error minifying {source_path}: {e}")
    failed_files.append((source_path, dest_path))

async def copy_file(source_path, dest_path, failed_files):
  try:
    os.makedirs(os.path.dirname(dest_path), exist_ok=True)
    async with aiofiles.open(source_path, 'rb') as fsrc:
      async with aiofiles.open(dest_path, 'wb') as fdst:
        await fdst.write(await fsrc.read())
  except Exception as e:
    print(f"Error copying {source_path}: {e}")
    failed_files.append((source_path, dest_path))

async def process_files(root, files, src_base, dest_base, failed_files):
  tasks = []
  for file in files:
    relative_path = os.path.relpath(os.path.join(root, file), src_base)
    dest_file_path = os.path.join(dest_base, relative_path)
    source_file_path = os.path.join(root, file)
    if file.endswith('.py'):
      tasks.append(minify_file(source_file_path, dest_file_path, failed_files))
    else:
      tasks.append(copy_file(source_file_path, dest_file_path, failed_files))
  await asyncio.gather(*tasks)

async def recreate_flask_app(source, destination):
  if os.path.exists(destination):
    shutil.rmtree(destination)
  os.makedirs(destination, exist_ok=True)

  tasks = []
  failed_files = []
  for root, dirs, files in os.walk(source):
    if should_skip(root):
      continue
    tasks.append(process_files(root, files, source, destination, failed_files))

  await asyncio.gather(*tasks)

  while failed_files:
    retry_files = failed_files
    failed_files = []
    retry_tasks = [minify_file(src, dest, failed_files) if src.endswith('.py') else copy_file(src, dest, failed_files) for src, dest in retry_files]
    await asyncio.gather(*retry_tasks)

if __name__ == "__main__":
  asyncio.run(recreate_flask_app(source_folder, destination_folder))
