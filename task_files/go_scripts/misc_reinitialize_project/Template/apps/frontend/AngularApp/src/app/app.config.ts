import { ApplicationConfig, ErrorHandler, importProvidersFrom, provideAppInitializer } from '@angular/core';
import { Router, TitleStrategy, provideRouter } from '@angular/router';
import { routes } from './app.routes';
import { provideClientHydration } from '@angular/platform-browser';
import { gsap } from "gsap";
import { ScrollToPlugin } from "gsap/ScrollToPlugin";
import { ENV } from '@env/environment';
import { HttpClient, HTTP_INTERCEPTORS, provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { HttpLoaderFactory, I18NSEOStrategy, waitFori18nextToLoad } from '@core/utility/i18n-utils';
import { TranslateModule, TranslateLoader } from '@ngx-translate/core';
import { GlobalErrorHandler } from '@shared/errorhandlers/global-error-handler';
import { NewAzureAccessTokenInterceptor } from '@shared/interceptors/new-azure-access-token.interceptor';
import { XsrfInterceptor } from '@shared/interceptors/xsrf.interceptor';




gsap.registerPlugin(ScrollToPlugin);



let routerProviders


if(ENV.app.shouldPerformInitialNavigate()){
  routerProviders = provideRouter(routes)
}
else{
  // routerProviders =provideRouter(routes, withDisabledInitialNavigation());
  routerProviders = provideRouter(routes)
}


export const appConfig: ApplicationConfig = {
  providers: [
    // provideExperimentalZonelessChangeDetection(),
    routerProviders,
    provideClientHydration(),
    provideHttpClient(withInterceptorsFromDi()),
    importProvidersFrom(
      TranslateModule.forRoot({
        defaultLanguage: 'en',
        loader: {
          provide: TranslateLoader,
          useFactory: HttpLoaderFactory,
          deps:[HttpClient]
        }
      })
    ),
    // @ts-ignore
    provideAppInitializer(waitFori18nextToLoad),

    {provide:HTTP_INTERCEPTORS,useClass:NewAzureAccessTokenInterceptor,multi:true},
    {provide:HTTP_INTERCEPTORS,useClass:XsrfInterceptor,multi:true},
    {provide:TitleStrategy,useClass:I18NSEOStrategy },




  ]
};
