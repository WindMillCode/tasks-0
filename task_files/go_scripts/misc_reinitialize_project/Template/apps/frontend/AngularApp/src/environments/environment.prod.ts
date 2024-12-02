import { silenceAllConsoleLogs, traverseClassAndRemoveUpdateAutomation } from "@core/utility/env-utils"
import { DevEnv } from "./environment.dev"


export let environment = {
  production: true
}
class ProdEnv extends DevEnv  {


  constructor(){
    super()
    this.type = "PROD"
    this.frontendURI0.url.host = "[PROJECT_NAME].com"
    this.backendURI0.url.host = "api.[PROJECT_NAME].com"
    ;[this.frontendURI0,this.backendURI0].forEach((uri)=>{
      uri.url.port = "443"
    })

    silenceAllConsoleLogs()
    traverseClassAndRemoveUpdateAutomation(this)




  }
}


export let ENV =   new ProdEnv()
