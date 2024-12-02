// angular
import { AfterViewInit, ChangeDetectorRef, Component, HostBinding, Renderer2, ViewContainerRef } from '@angular/core';
import { Router, RouterOutlet } from '@angular/router';

// rxjs
import { catchError, delay, fromEvent, merge, of, Subject, takeUntil, tap, timer } from 'rxjs';

// services
import { BaseService } from '@core/base/base.service';

// misc
import { ENV } from '@env/environment';
import {  UtilityService } from '@core/utility/utility.service';

// wml-components
import { HttpClient } from '@angular/common/http';
import { WMLNotifyOneService } from '@windmillcode/angular-wml-notify';
import { NavService } from '@shared/services/nav/nav.service';
import { generateClassPrefix } from '@windmillcode/wml-components-base';
import { changeAllButtonTypeAttributesToButton, documentQuerySelector, documentQuerySelectorAll } from '@core/utility/common-utils';
import { SharedModule } from '@shared/shared.module';
import { SITE_OFFLINE_ENUM } from '@core/utility/constants';
import { SpecificService } from './core/specific/specific.service';
import {gsap} from "gsap";

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
    imports: [
        RouterOutlet,
        SharedModule
    ],
    providers: []
    // changeDetection: ChangeDetectionStrategy.OnPush
})
export class AppComponent implements AfterViewInit {
  constructor(
    public baseService: BaseService,
    public utilService: UtilityService,
    public cdref: ChangeDetectorRef,
    public vcf: ViewContainerRef,
    public router:Router,
    public http:HttpClient,
    public WMLNotifyOneService:WMLNotifyOneService,
    public navService:NavService,
    public renderer2:Renderer2,
    public specificService:SpecificService
  ) {
    this.listenForChangesOutSideChangeDetection().subscribe()
  }

  classPrefix = generateClassPrefix(ENV.classPrefix.app)
  @HostBinding('class') myClass: string = this.classPrefix(`View`);
  ngUnsub = new Subject<void>()


  listenForChangesOutSideChangeDetection = ()=>{
    return merge(
      this.baseService.popupProps.togglePopupSubj,
      this.baseService.toggleOverlayLoadingSubj,
    )
    .pipe(
      takeUntil(this.ngUnsub),
      tap(()=>{
        this.cdref.detectChanges()
      })
    )

  }

  removeAngularIdentifiers() {
    if (!["DEV","TEST"].includes(ENV.type)) {
      this.vcf.element.nativeElement.removeAttribute('ng-version');
    }
  }

  handleSiteNavigation = ()=>{
    if (ENV.app.siteOffline === SITE_OFFLINE_ENUM.TRUE) {

      this.router.navigateByUrl(ENV.nav.urls.siteOffline);
      return false
    } else if (this.utilService.getWindow().location.pathname === ENV.nav.urls.siteOffline) {
      this.router.navigateByUrl(ENV.nav.urls.home);
      return true
    }
    return true
  }

  checkIfBackendIsWorking() {
    return this.http
      .get(ENV.app.backendHealthCheck())
      .pipe(
        takeUntil(this.ngUnsub),
        catchError((err) =>{return of(null)}
          // throwError(() => {
          //   // this.router.navigateByUrl(ENV.nav.urls.siteOffline);
          //   // return new CantReachBackendError(err);
          // })
        )
      );
  }




  doMiscConfigs() {
    this.removeAngularIdentifiers();
    let myContinue = this.handleSiteNavigation()
    if(myContinue){
      this.baseService.appCdRef = this.cdref;
      ENV.nav.urls.initialURL = window.location.pathname;
      if(!['DEV','TEST','DOCKER_DEV'].includes(ENV.type)){
        this.checkIfBackendIsWorking().subscribe();
      }
    }
    return of()
    .pipe(
      takeUntil(this.ngUnsub),

      tap(()=>{

        this.automate()
      }),


    )


  }



  automate =()=>{
    // this.navService.mobileNavOne.open()
    if(["DOCKER_DEV","DEV"].includes(ENV.type)){
    fromEvent(window,"load")
    .pipe(
      takeUntil(this.ngUnsub),
      delay(2000),
      tap(()=>{

        gsap.to(
          document.querySelector(".LandingZeroLayoutMainPod"),
          {
            duration:1,
            scrollTo:"#preview"
          }
        )
      }),

    )
    .subscribe()
    }
  }

  ngOnInit() {
    changeAllButtonTypeAttributesToButton()
    this.doMiscConfigs().subscribe()
  }

  ngAfterViewInit (){

    if(this.utilService.getWindow().navigator.userAgent !=="Puppeteer"  ){


      timer(1000)
      .pipe(
        takeUntil(this.ngUnsub),
        tap(()=>{
          let appRoots = documentQuerySelectorAll("app-root");
          let appRootToRemove = documentQuerySelector("app-root:not(.AppSSGFrame0)")
          let appRootToUpdateClass = documentQuerySelector("app-root:nth-of-type(1)")
          if (appRootToRemove && appRoots.length > 1 && ["PREVIEW","PROD"].includes(ENV.type)) {
            this.cdref.detach()
            appRootToRemove.remove();
            this.cdref.reattach()
          }
          appRootToUpdateClass.classList.add("AppSSGInit")
          this.cdref.detectChanges()
        })
      )
      .subscribe()

    }
  }


  ngOnDestroy() {
    this.ngUnsub.next()
    this.ngUnsub.complete()
  }

}
