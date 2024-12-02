import { SITE_OFFLINE_ENUM } from '@core/utility/constants';
import { EnvPlatformType, traverseClassAndRemoveUpdateAutomation } from '@core/utility/env-utils';
import { makeTitleCase } from '@core/utility/string-utils';

import { WMLUri } from '@windmillcode/wml-components-base';


export let environment = {
  production: false,
};


export class DevEnv {
  // this must be a union type for import purposes
  type: 'DEV' | 'DOCKER_DEV' | 'PREVIEW' | 'PROD' | 'TEST' = 'DEV';
  platformType = EnvPlatformType.WEB
  endpointMsgCodes = {
    success: 'OK',
    error: 'ERROR',
    respViaSocketio: 'RESPONSE_VIA_SOCKETIO',
  };

  app = (() => {
    let app = {

      shouldPerformInitialNavigate:()=>{
        return navigator.userAgent ==="Puppeteer" || !["PREVIEW","PROD"].includes(this.type)
      },
      backendHealthCheck: () => this.backendURI0.fqdn + '/healthz/',
      siteOffline: SITE_OFFLINE_ENUM.FALSE,
      originalConsoleObj: [],
      keyBoardDebounceTimes:[500],
      isRemote:false //remember when setting to true you have to run the app in http because the emulator rejects your certificate as its not in the cert store of the remote device // TODO just know when deploying always set isRemote to false

      // dev additions



    };
    return app;
  })();

  backendURI0 = new WMLUri({
    host:"example.com",
    port:[Flask_Run_0]
  })
  frontendURI0 = new WMLUri({
    host:"example.com",
    port:[Angular_Run_0]
  })

  classPrefix = {
    app: 'App'
  };

  idPrefix: { [k: string]: string } = {

    confirmDialogZero: 'ConfirmDialogZero_',
    logoImgZero: 'LogoImgZero_',

    landingZeroLayout:"LandingZeroLayout_",
    powerBiZeroPage:"PowerBiZeroPage_",


  };

  nav = (() => {
    let nav = {
      urls: {

        home: '/',
        homeAlt: '/home',
        siteOffline: '/site-offline',
        notFound:"/404",
        initialURL: '/',

        landingZeroLayout:"/landing",

        powerBiZeroPage:"/power-bi",
        // DEV ADDITIONS

      },
    };

    let idPrefixes = Object.entries(nav.urls).map(([key, val]) => {
      return [key, makeTitleCase(key) + '_'];
    });
    this.idPrefix = {
      ...Object.fromEntries(idPrefixes),
      ...this.idPrefix,
    };
    return nav;
  })();



  errorCodes = {};


  constructor() {
    this.app.originalConsoleObj = Object.entries(console);
    traverseClassAndRemoveUpdateAutomation(this)
    if(this.app.isRemote){

    }
  }
}

export let ENV = new DevEnv();
