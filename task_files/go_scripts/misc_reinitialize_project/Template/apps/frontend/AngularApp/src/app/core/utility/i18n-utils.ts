import {inject, Injectable} from '@angular/core';
import { RouterStateSnapshot, TitleStrategy} from '@angular/router';
import {TranslateService} from '@ngx-translate/core';
import {Meta, Title} from '@angular/platform-browser';
import { HttpClient } from '@angular/common/http';
import { TranslateHttpLoader } from '@ngx-translate/http-loader';
import { tap, zip } from 'rxjs';
import { ENV } from '@env/environment';


@Injectable()
export class I18NSEOStrategy extends TitleStrategy {
  constructor(
    public translateService: TranslateService,
    public  title: Title,
    public meta: Meta
  ) {
    super();
  }

  updateCanonicalUrl(url: string) {
    const canonicalUrl = document.createElement('link');
    canonicalUrl.setAttribute('rel', 'canonical');
    // TODO update the nagivation logic so the cannonical is valid for the root page
    if(url.includes("landing")){
      url = "/"
    }
    let origin = window.location.origin
    if(window.navigator.userAgent ==="Puppeteer" ){
      // TODO should look like 'https://example.com:4203' sumbit a bug as report as necessarey
      origin =ENV.frontendURI0.fqdn
    }
    canonicalUrl.setAttribute('href', `${origin}${url}`);
    const existingCanonicals = document.querySelectorAll('link[rel="canonical"]');
    existingCanonicals.forEach(canonical => {
      canonical.parentNode.removeChild(canonical);
    });
    document.head.appendChild(canonicalUrl);
  }

  updateJSONLD = ([translatedTitle,translatedDesc,translatedKeywords])=>{
    const jsonLDTags = document.querySelectorAll('script[type="application/ld+json"]');
    jsonLDTags.forEach(tag => {
      tag.parentNode.removeChild(tag);
    });

    const jsonLDScript = document.createElement('script');
    jsonLDScript.type = 'application/ld+json';
    jsonLDScript.text = JSON.stringify({
      "@context": "https://schema.org",
      "@type": "WebPage",
      "name": translatedTitle,
      "description": translatedDesc,
      "keywords": translatedKeywords.split(',').map(keyword => keyword.trim())
    });
    document.head.appendChild(jsonLDScript);
  }

  updateGeneralMetaTags = ([translatedTitle,translatedDesc,translatedKeywords])=>{

    this.title.setTitle(translatedTitle);

    const metaTags = document.querySelectorAll('meta[name="description"],meta[name="keywords"]');
    metaTags.forEach(canonical => {
      canonical.parentNode.removeChild(canonical);
    });
    this.meta.addTags(["description","keywords"].map((name,index0)=>{
      return{
        name,content:[translatedDesc,translatedKeywords][index0]
      }
    }))
  }


  updateSEO(snapshot: RouterStateSnapshot) {
    const title = this.buildTitle(snapshot);
    let {url} = snapshot
    return zip(
      this.translateService.get((title ?? "nav.pageInfo.default") + ".title"),
      this.translateService.get((title ?? "nav.pageInfo.default") + ".description"),
      this.translateService.get((title ?? "nav.pageInfo.default") + ".keywords")
    )
      .pipe(
        tap(() => {
          this.updateCanonicalUrl(url);
        }),
        tap(this.updateJSONLD),
        tap(this.updateGeneralMetaTags)
      );
  }

  override updateTitle(snapshot: RouterStateSnapshot): void {

    this.updateSEO(snapshot)
    .subscribe()
  }
}




export function HttpLoaderFactory(http: HttpClient) {
  return new TranslateHttpLoader(http);
}

export function waitFori18nextToLoad() {
  const translateService = inject(TranslateService);
  return () => {
    return translateService.use('en')
  }
}


