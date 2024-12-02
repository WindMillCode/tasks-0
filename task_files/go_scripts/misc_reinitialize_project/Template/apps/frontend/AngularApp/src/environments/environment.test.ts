import { traverseClassAndRemoveUpdateAutomation } from "@core/utility/env-utils"
import { DevEnv } from "./environment.dev"


export let environment = {
  production: true
}
export class TestEnv extends DevEnv  {


  constructor(){
    super()
    this.type = "TEST"
    traverseClassAndRemoveUpdateAutomation(this,[])

  }
}


export let ENV =   new TestEnv()
