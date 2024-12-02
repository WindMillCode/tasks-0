import { silenceAllConsoleLogs, traverseClassAndRemoveUpdateAutomation } from "@core/utility/env-utils"
import { DevEnv } from "./environment.dev"


export let environment = {
  production: true
}
class DockerDevEnv extends DevEnv  {


  constructor(){
    super()
    this.type = "DOCKER_DEV"
    this.backendURI0.url.host = "localhost"
    this.backendURI0.url.port = 5000
    silenceAllConsoleLogs()
    traverseClassAndRemoveUpdateAutomation(this)
  }
}


export let ENV =   new DockerDevEnv()
