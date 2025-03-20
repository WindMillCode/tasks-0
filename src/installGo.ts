import { mkdirSync } from "fs";
import * as os from "os";
import * as path from "path"
import {  notifyDeveloper } from './functions';
import * as fs from 'fs';
import { exec } from 'child_process';
import { promisify } from 'util';
import * as vscode from 'vscode';
const semver = require('semver');

let extensionDesiredVersion = "1.24.1"


let checkGoInstalledInExtension =  (desiredVersion=extensionDesiredVersion)=>{
  return new Promise(async(resolve,rej)=>{
    let executable = "go"
    exec(`${executable} version`,async  (error, stdout, stderr) => {
      if (error) {
        resolve(false);
      } else {
        let versionMatch = stdout.match(/go(\d+\.\d+\.\d+)/);
        let installedVersion = versionMatch ? versionMatch[1] : null;
        notifyDeveloper(installedVersion)
        notifyDeveloper(desiredVersion)
        if(!semver.gt(desiredVersion,installedVersion)){
          notifyDeveloper(null,`Go is installed and the correct version is in the extension ${stdout.trim()}`);
          resolve(executable);
        }
        else{
          notifyDeveloper(null,`It seems the correct version is not installed on the system.${stdout.trim()}`);
          resolve(false)
        }
      }
    });
  })
}




async function addToPath(directory:string) {
  let platform:Partial<NodeJS.Platform> = os.platform()
  // @ts-ignore
  let actions ={
    "win32":{
      "env_setter":"setx PATH "
    },
    "darwin":{
      "env_setter":"export PATH="
    },
    "linux":{
      "env_setter":"export PATH="
    }
  }[platform]
  if(process.env.PATH?.includes(directory)){
    notifyDeveloper(null,"go executable is already on the path")
    concatPath(platform, directory);
    changePermission(directory)
    return Promise.resolve(true)
  }
  let cleanedPath = process.env.PATH?.split(path.delimiter)
  .filter((path)=>!path.includes("windmillcode"))
  .join(path.delimiter)

  process.env.PATH = `${directory}${path.delimiter}${cleanedPath}`;
  concatPath(platform, directory);
  changePermission(directory)
  return new Promise((res,rej)=>{
    exec(`${actions?.env_setter}"${process.env.PATH}"`,(err,stdout,stderr)=>{
      if(err){
        notifyDeveloper(null,err.stack)
        res(false)
      }
      notifyDeveloper(null,"added to path sucessfully")
      res(true)
    })
  })

}


function concatPath(platform: string, directory: string) {
  if (platform === "darwin") {
    let execStart = "export PATH=$PATH"+path.delimiter;
    let finalExecMac = execStart +directory;
    editZshrcFile(finalExecMac);
  }
}

function changePermission(directory: string){
  let fileEnd = "windmillcode_go"
  let finalFile = path.join(fileEnd)
  fs.chmod(finalFile,0o777,() =>{
    notifyDeveloper(null,`The permissions for ${finalFile} changed`);
  })
}

async function copyFile(sourcePath:string, destinationPath:string) {
  const readFileAsync = promisify(fs.readFile);
  const writeFileAsync = promisify(fs.writeFile);

  try {
    const data = await readFileAsync(sourcePath);
    await writeFileAsync(destinationPath, data);
    console.log('File copied successfully.');
  } catch (error) {
    console.error('Error copying the file:', error);
  }
}



export let installGo = async (extensionRoot:string,goVersion=extensionDesiredVersion,) => {
  // Change these values as needed
  let installLocation = path.normalize(extensionRoot+"/task_files")
  notifyDeveloper(installLocation)

  try {
    mkdirSync(installLocation, { recursive: true });
  } catch (err) {
    notifyDeveloper("Error creating installation directory:");
    process.exit(1);
  }

  // Determine the platform and letruct platform-specific commands and paths
  let platform = os.platform();


  // Download and install Go

  // @ts-ignore
  let executable:boolean |"go"  = await checkGoInstalledInExtension()



  notifyDeveloper(null, ` Executable ${executable}`)
  if(executable === false){
    let message = `Golang is not installed on the system. Please install golang with a version greater than or equal to ${goVersion}`
    vscode.window.showErrorMessage(message)
    throw message
  }
  else{
    return  {
      executable:"go",
      alreadyInstalled:true
    }
  }




}

function editZshrcFile(newContent: string): void {
  const homeDir = process.env.HOME || process.env.USERPROFILE;

  if (!homeDir) {
    console.error("Home directory not found.");
    return;
  }

  const zshrcPath = path.join(homeDir, '.zshrc');

  fs.readFile(zshrcPath, 'utf8', (err, data) => {
    if (err) {
      console.error("Error reading .zshrc:", err);
      return;
    }

    let modifiedContent = data + '\n' + newContent;
    modifiedContent = modifiedContent.split('\n')
    .filter((path)=>!path.includes("windmillcode"))
    .join('\n')

    fs.writeFile(zshrcPath, modifiedContent, 'utf8', (err) => {
      if (err) {
        console.error("Error writing to .zshrc:", err);
        return;
      }

      console.log(".zshrc has been successfully updated.");
    });
  });
}
