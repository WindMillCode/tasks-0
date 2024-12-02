import { Injectable } from '@angular/core';
import { ENV } from '@env/environment';
import { BaseService } from '@core/base/base.service';
import { WMLNotifyOneService } from '@windmillcode/angular-wml-notify';
import { WMLCustomComponent, WMLImage } from '@windmillcode/wml-components-base';
import { UtilityService } from '@core/utility/utility.service';
import { GenerateFieldProps, createWMLFormZeroPropsField } from '@core/utility/form-utils';

import { WMLFileUploadZeroProps, WMLFileUploadZeroComponent } from '@windmillcode/angular-wml-file-manager';
import { WMLChipsZeroProps, WMLChipsZeroComponent } from '@windmillcode/angular-wml-chips';


import { defer } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SpecificService  {
  constructor(
    public utilService:UtilityService,
    public WMLNotifyOneService:WMLNotifyOneService,
    public baseService:BaseService
  ){}

  getLogoImg = ()=>new WMLImage({
    src:"assets/media/app/logo-no-bg.png",
    alt:"global.logoImgAlt",
    routerLink:ENV.nav.urls.home
  })
  logoImg = this.getLogoImg()



  _manageWebStorage(webStorage:Storage,thisWebStorage:any,predicate:Function) {
    let myWebStorage = webStorage.getItem(ENV.classPrefix.app);
    myWebStorage = JSON.parse(myWebStorage);
    Object.assign(thisWebStorage, myWebStorage);
    predicate()
    webStorage.setItem(ENV.classPrefix.app, JSON.stringify(thisWebStorage))
  }




  getWakeLock = ()=>{
    return defer(async ()=>{
      try{
        return await navigator.wakeLock.request('screen');
      } catch (err){
        return null
      }
    })
  }



  createField=(props =new GenerateFieldProps<any>(),cpnt:WMLCustomComponent["cpnt"])=>{
    let {
      fieldCustomProps
    } = props
    fieldCustomProps = fieldCustomProps
    let wmlField = createWMLFormZeroPropsField(props)
    wmlField.field.custom.cpnt = cpnt
    wmlField.field.custom.props  =fieldCustomProps

    return wmlField
  }



  createWMLFileUploadField = (props= new GenerateFieldProps())=>{
    props.fieldCustomProps ??= new WMLFileUploadZeroProps()
    return this.createField(props,WMLFileUploadZeroComponent)
  }

  createWMLChipsZeroField =(props =new GenerateFieldProps())=>{
    props.fieldCustomProps ??= new WMLChipsZeroProps()
    return this.createField(props,WMLChipsZeroComponent)
  }





}
