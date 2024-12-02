#!/bin/sh

/usr/sbin/sshd
crond
git init .
git remote add origin git@github.com:[ORGANIZATION_NAME_CAPITAL]/[PROJECT_NAME_CAPITAL].git
git config core.sparsecheckout true
echo "apps/backend/FlaskApp*" >> .git/info/sparse-checkout


if [ "$DEBUG_CODE" == "TRUE" ]; then
  git pull origin [ORGANIZATION_NAME]-winners --depth=1
fi
if [ "$FLASK_BACKEND_ENV" == "PREVIEW" ]; then
  if [ "$DEBUG_CODE" != "TRUE" ]; then
    git pull origin dev --depth=1
  fi

  mv apps/backend/FlaskApp/* ./
  mv /[PROJECT_NAME]-preview-gae.json ./
  mv /[PROJECT_NAME]-preview-client_secret.json ./
  mv /[PROJECT_NAME]-preview-ca-2021.crt ./
else
  if [ "$DEBUG_CODE" != "TRUE" ]; then
    git pull origin main --depth=1
  fi
  mv apps/backend/FlaskApp/* ./
  mv /[PROJECT_NAME]-prod-gae.json ./
  mv /[PROJECT_NAME]-prod-client_secret.json ./
  mv /[PROJECT_NAME]-prod-ca-2021.crt ./
fi
mv /[PROJECT_NAME]-apple-iap-key.p8 ./
mv /apple_root_certs ./my_resources

pip install -r linux-requirements.txt
pip install --pre yt-dlp[default]

if [ "$DEBUG_APP" == "TRUE" ]; then
  python app.py
else
  gunicorn -k gevent -w 1 app:app  --log-file app.log --log-level debug --access-logfile ./access.log --access-logformat "%(h)s %(l)s %(u)s %(t)s '%(r)s' %(s)s %(b)s '%(f)s' '%(a)s'"
fi
