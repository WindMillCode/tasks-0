import * as path from "path"
import {  notifyDeveloper } from './functions';
import * as fs from 'fs';
import { exec } from 'child_process';
import * as vscode from 'vscode';
const semver = require('semver');

let extensionDesiredVersion = "1.24.2"


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
        }
        else{
          notifyDeveloper(null,`It seems the correct version is not installed on the system.${stdout.trim()}. You can try to use the plugin but if there are any issues please upgrade golang to ${desiredVersion}`);
        }
        resolve(executable);
      }
    });
  })
}




export let checkForGolang = async (goVersion=extensionDesiredVersion) => {


  // @ts-ignore
  let executable:boolean |"go"  = await checkGoInstalledInExtension()



  notifyDeveloper(null, ` Executable ${executable}`)
  if(executable === false){
    let message = `Golang is not installed on the system. Please install golang with a version greater than or equal to ${goVersion}`
    vscode.window.showErrorMessage(message)
  }
  return  {
    executable:"go",
    alreadyInstalled:true
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
