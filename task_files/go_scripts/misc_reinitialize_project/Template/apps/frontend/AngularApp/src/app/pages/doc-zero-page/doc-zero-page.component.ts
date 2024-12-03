import { ChangeDetectionStrategy, ChangeDetectorRef, Component, HostBinding } from '@angular/core';
import { UtilityService } from '@app/core/utility/utility.service';
import { BaseService } from '@core/base/base.service';
import { Subject } from 'rxjs';
import { WMLUIProperty, generateClassPrefix, generateIdPrefix } from '@windmillcode/wml-components-base';
import { ENV } from '@env/environment';
import enTranslations from "src/assets/i18n/en.json";
import { SharedModule } from '@shared/shared.module';


type DocZeroPageType = "text"|"section"|"list" |"value"| "bold_text";
@Component({
    selector: 'doc-zero-page',
    templateUrl: './doc-zero-page.component.html',
    styleUrls: ['./doc-zero-page.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [
        SharedModule
    ]
})
export class DocZeroPageComponent  {
  constructor(
    public cdref:ChangeDetectorRef,
    public utilService:UtilityService,
    public baseService:BaseService
  ) { }

  classPrefix = generateClassPrefix('DocZeroPage')
  idPrefix = generateIdPrefix(ENV.idPrefix.DocZeroPage)
  @HostBinding('class') myClass: string = this.classPrefix(`View`);
  @HostBinding('id') myId:string = this.idPrefix()
  ngUnsub= new Subject<void>()



  optionsObj = {
    [ENV.nav.urls.privacyPolicyPage]:{
      title:"PrivacyPolicy.title",
      sections:this.createSections({
        sections:enTranslations.PrivacyPolicy.sections,
        i18nSectionString:"PrivacyPolicy.sections"
      })
    },
    [ENV.nav.urls.termsOfServicePage]:{
      title:"TermsOfService.title",
      sections:this.createSections({
        sections:enTranslations.TermsOfService.sections,
        i18nSectionString:"TermsOfService.sections"
      })
    }
  }[this.utilService.router.url]
  sections = this.optionsObj.sections

  createSections(props){
    let {sections,i18nSectionString} = props
    return sections
    .map((sectionText,index0)=>{
      let type = Object.keys(sectionText)[0]
      let i18nText:any = this.transformSectionText(sectionText,`${i18nSectionString}.${index0}`,type)

      return new WMLUIProperty<DocZeroPageType>({
        text:i18nText,
        type:type?? "text"
      })
    })
  }

  transformSectionText<T>(
    // TODO fix typing
    item: any,
    path: string,
    type:string
  ): any {
    // Recursive case: item is a array
    if (item?.list) {
      return item.list.map((nestedItem, index) => {
        // Construct the new path for this nested item
        const newPath = `${path}.${Object.keys(item)[0]}.${index}`;
        // Recursively traverse and transform this nested item
        return this.transformSectionText(nestedItem, newPath,Object.keys(nestedItem)[0]);
      });
    }
    // Base case: item is an object
    else  {
      return `${path}.${type}`;
    }

  }



  ngOnDestroy(){
    this.ngUnsub.next();
    this.ngUnsub.complete()
  }
}


