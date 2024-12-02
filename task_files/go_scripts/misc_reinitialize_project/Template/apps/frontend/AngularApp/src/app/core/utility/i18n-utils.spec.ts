import { TranslateService } from '@ngx-translate/core';
import { of } from 'rxjs';
import { wmlTestUtils } from './test-utils/test-utils';
import { HttpLoaderFactory, I18NPageTitleStrategy, waitFori18nextToLoad } from './i18n-utils';
import { fakeAsync, discardPeriodicTasks } from '@angular/core/testing';
import { flush } from '@sentry/angular-ivy';
import { HttpClient } from '@angular/common/http';
import { TranslateHttpLoader } from '@ngx-translate/http-loader';
import { Title } from '@angular/platform-browser';
import { RouterStateSnapshot } from '@angular/router';
describe("I18NUtils",()=>{

  describe('waitFori18nextToLoad', () => {
    let translateService: TranslateService;

    beforeEach(() => {
      // Initialize TranslateService as needed
      translateService = wmlTestUtils.mock.mockTranslateService
      spyOn(translateService,"use").and.callThrough()
    });

    it(` when called |
      under normal conditions |
      does the required action `,fakeAsync(()=>{
      // arrange
      let fn = (val?)=>{
        // assert
        expect(translateService.use).toHaveBeenCalledWith('en')
        discardPeriodicTasks()
      }
      const waitFori18next = waitFori18nextToLoad(translateService);

       // act
       waitFori18next()
      .subscribe({next:fn,complete:fn})
      // @ts-ignore
      translateService.useResult.next("")
      flush()

      }))

  });

  describe('HttpLoaderFactory', () => {
    let http: HttpClient;

    beforeEach(() => {
      // Initialize HttpClient as needed
      http = jasmine.createSpyObj('HttpClient', ['get']);
    });

    it('should return a TranslateHttpLoader instance', () => {
      // Arrange
      const result = HttpLoaderFactory(http);


      // Assert
      expect(result).toBeInstanceOf(TranslateHttpLoader);
    });
  });


describe('I18NPageTitleStrategy', () => {
  let i18nPageTitleStrategy: I18NPageTitleStrategy;
  let translateService: TranslateService;
  let titleService: Title;

  beforeEach(() => {
    translateService = jasmine.createSpyObj('TranslateService', ['get']);
    titleService = jasmine.createSpyObj('Title', ['setTitle']);

    i18nPageTitleStrategy = new I18NPageTitleStrategy(translateService, titleService);
  });

  describe('updateTitle', () => {
    it('should update the title with translated value when called', () => {
      // Arrange
      const snapshot: RouterStateSnapshot = {} as RouterStateSnapshot;
      const translatedTitle = 'Translated Title';
      (translateService.get as jasmine.Spy).and.returnValue(of(translatedTitle));

      // Act
      i18nPageTitleStrategy.updateTitle(snapshot);

      // Assert
      expect(translateService.get).toHaveBeenCalledWith(jasmine.any(String));
      expect(titleService.setTitle).toHaveBeenCalledWith(translatedTitle);
    });

    it('should use default title if translation is not available', () => {
      // Arrange
      const snapshot: RouterStateSnapshot = {} as RouterStateSnapshot;
      (translateService.get as jasmine.Spy).and.returnValue(of(undefined));
      const defaultTitle = 'Default Title';

      // Act
      i18nPageTitleStrategy.updateTitle(snapshot);

      // Assert
      expect(translateService.get).toHaveBeenCalledWith(jasmine.any(String));
      expect(titleService.setTitle).toHaveBeenCalledWith(undefined);
    });
  });
});

})
