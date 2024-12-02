import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UtilityService } from '@core/utility/utility.service';
import { WMLInfiniteDropdownZeroProps } from '@windmillcode/angular-wml-infinite-dropdown';
import { BaseService } from '@core/base/base.service';
import { SpecificService } from '@core/specific/specific.service';
import enTranslations from "src/assets/i18n/en.json";
import { replaceValuesWithPaths, WMLRoute } from '@windmillcode/wml-components-base';
import { ENV } from '@env/environment';

@Injectable({
  providedIn: 'root',
})
export class NavService {
  constructor(
    public http: HttpClient,
    public utilService: UtilityService,
    public baseService: BaseService,
    public specificService:SpecificService
  ) {}
  // same structure for mobileNav you should seperate desktop and mobile nav
  createDesktopNavZero = () => {
    let i18nDropdown: any = replaceValuesWithPaths(
      enTranslations.NavZero.desktopOptions,
      'NavZero.desktopOptions.',
      ({key,value,path,defaultAssignment})=>{
        return new WMLRoute({
          text:defaultAssignment
        })
      }
    );
    return i18nDropdown
  }

  setActionsForDesktopNavZeroOptions = (options) => {
    options.forEach((link:WMLRoute,index0)=>{
      link.routerLink = ENV.nav.urls[["home"][index0]]
    })
  };

  addIdsToDesktopNavZeroOptions = (options) => {

    let ids = [
      ["homeLink"]
    ]
    options.forEach((link,index0)=>{
      link.id = ENV.idPrefix.nav+ids[index0]
    })

  };



  desktopNavZero = (() => {
    let mainDropdowns = this.createDesktopNavZero();
    this.setActionsForDesktopNavZeroOptions(mainDropdowns)
    this.addIdsToDesktopNavZeroOptions(mainDropdowns);

    return mainDropdowns;
  })();




  getOptions(dropdowns: WMLInfiniteDropdownZeroProps[],makeResultSet= false) {
    return dropdowns
      .map((dropdown) => {
        if(makeResultSet){
          return {
            dropdown,
            option:dropdown.options[0]
          }
        }
        return dropdown.options[0]
      }) as any;
  }
}
