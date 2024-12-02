// testing
import { wmlTestUtils } from '@core/utility/test-utils/test-utils';
import { resetImports } from '@core/utility/test-utils/mock-imports';
import { resetProviders } from '@core/utility/test-utils/mock-providers';
import { resetDeclarations } from '@core/utility/test-utils/mock-declarations';
import { TestBed } from '@angular/core/testing';

// services
import { UtilityService } from '@core/utility/utility.service';

import { NavService } from './nav.service';
import { WMLInfiniteDropdownZeroProps } from '@windmillcode/angular-wml-infinite-dropdown';
import { ENV } from '@env/environment';

let wmlComponentsBase = require('@windmillcode/wml-components-base')
import { gsap } from "gsap";

let commonUtils = require("@core/utility/common-utils")
import enTranslations from "src/assets/i18n/en.json";
import { makeLowerCase } from '@core/utility/string-utils';

describe('NavService', () => {
  let service: NavService;
  let utilService:UtilityService
  let htmlElement = document.createElement("div")

  beforeEach(() => {
    resetImports()
    resetProviders(NavService);
    resetDeclarations()
    service = wmlTestUtils.configureTestingModuleForServices(NavService)
    utilService =TestBed.inject(UtilityService)


    wmlTestUtils
      .spyOnForES6Imports(wmlComponentsBase,'replaceValuesWithPaths')
      .and.callFake((val)=>val)



    spyOn(gsap,"to")
    spyOn(commonUtils,"documentQuerySelector").and.returnValue(htmlElement)
    spyOn(service.utilService.router,"navigateByUrl")
    spyOn(service.utilService.router,"navigate")
    spyOn(wmlTestUtils.mock.window,"open")
    spyOn(service.accountsService,"signOutViaFirebase").and.callThrough()
    spyOn(service.jobsService.jobPanelItem, "open").and.callThrough();
    spyOn(service.utilService, "changeLanguage").and.callThrough();

  });

  describe('init', () => {
    it('should create', () => {
      expect(service).toBeTruthy();
    });

    it('should have all values initialized properly', () => {
      // Expectations for class property primitives

    });

    it('should have all properties be the correct class instance', () => {
      // Expectations for class property instances
      service.desktopNavOne.forEach((dropdown) => {
        expect(dropdown).toBeInstanceOf(WMLInfiniteDropdownZeroProps);
      });
      // Add more property instance expectations as needed
    });
  });

  describe('createDesktopNavZero', () => {
    it('should create desktop navigation zero with the correct structure', () => {
      // Arrange & Act
      const navZero = service.createDesktopNavZero();
      // Assert
      expect(navZero).toBeInstanceOf(Array);
      navZero.forEach((dropdown) => {
        expect(dropdown).toBeInstanceOf(WMLInfiniteDropdownZeroProps);
      });
    });
  });


  describe('setActionsForDesktopNavZeroOptions', () => {
    it('should set click actions for desktop navigation options', () => {
      // Arrange
      const dropdowns = service.createDesktopNavZero();

      // Act
      service.setActionsForDesktopNavZeroOptions(dropdowns);
      // Assert
      dropdowns.forEach((dropdown, index) => {
        const option = dropdown.options[0];
        option.click();
        if (index < 2) {
          expect(service.utilService.router.navigate).toHaveBeenCalledWith([jasmine.any(String)]);
        } else if (index === 2) {
          expect(service.jobsService.jobPanelItem.open).toHaveBeenCalled();
        }
      });
    });
  });

  describe('addIdsToDesktopNavZeroOptions', () => {
    it('should add IDs to desktop navigation zero options', () => {
      // Arrange
      const dropdowns = service.createDesktopNavZero();
      // Act
      service.addIdsToDesktopNavZeroOptions(dropdowns);
      // Assert
      const expectedIds = ['managerOption', 'hubOption', 'trackerOption'];
      dropdowns.forEach((dropdown, index) => {
        expect(dropdown.id).toEqual(ENV.idPrefix.nav + expectedIds[index]);
      });
    });
  });


  describe('createDesktopNavOne', () => {
    it('should create desktop navigation one with the correct structure', () => {
      // Arrange & Act
      const navOne = service.createDesktopNavOne();
      // Assert
      expect(navOne).toBeInstanceOf(Array);
      navOne.forEach((dropdown) => {
        expect(dropdown).toBeInstanceOf(WMLInfiniteDropdownZeroProps);
      });
    });
  });

  describe('setActionsForDesktopNavOneOptions', () => {
    it('should set click actions for desktop navigation one options', () => {
      // Arrange
      const dropdowns = service.createDesktopNavOne();
      // Act
      service.setActionsForDesktopNavOneOptions(dropdowns);
      // Assert
      dropdowns.forEach((dropdown, index) => {
        const option = dropdown.options[0];
        option.click();
        if (index === 0) {
          expect(service.utilService.router.navigate).toHaveBeenCalledWith([ENV.nav.urls.connectionHub]);
        } else if (index === 4) {
          expect(service.accountsService.signOutViaFirebase).toHaveBeenCalled();
        } else {
          const ctrnIndex = {1: 0, 2: 2, 3: 3}[index];
          // @ts-ignore
          expect(gsap.to).toHaveBeenCalledWith(jasmine.anything(), jasmine.objectContaining({
            scrollTo: jasmine.objectContaining({
              y: jasmine.anything(),
              offsetY: jasmine.anything()
            })
          }));
        }
      });
    });
  });

  describe('addIdsToDesktopNavOneOptions', () => {
    it('should add IDs to desktop navigation one options', () => {
      // Arrange
      const dropdowns = service.createDesktopNavOne();
      // Act
      service.addIdsToDesktopNavOneOptions(dropdowns);
      // Assert
      const expectedIds = ['dashboardOption', 'profileOption', 'billingOption', 'dataOption', 'signOutOption'];
      dropdowns.forEach((dropdown, index) => {
        expect(dropdown.id).toEqual(ENV.idPrefix.nav + expectedIds[index]);
      });
    });
  });


  describe('createMobileNavZero', () => {
    it('should create mobile navigation zero with the correct structure', () => {
      // Arrange & Act
      const navZero = service.createMobileNavZero();
      // Assert
      expect(navZero).toBeInstanceOf(Array);
      navZero.forEach((dropdown) => {
        expect(dropdown).toBeInstanceOf(WMLInfiniteDropdownZeroProps);
      })
      navZero.slice(1).forEach((dropdown) => {
        expect(dropdown).toBeInstanceOf(WMLInfiniteDropdownZeroProps);
      });
    });
  });

  describe('setActionsForMobileNavZeroOptions', () => {
    it('should set click actions for mobile navigation zero options', () => {
      // Arrange
      const dropdowns = service.createMobileNavZero();
      // Act
      service.setActionsForMobileNavZeroOptions(dropdowns);
      // Assert
      dropdowns.slice(2).forEach((dropdown, index) => {
        const option = dropdown.options[0];
        option.click();
        if (index < 6) {
          expect(service.utilService.router.navigate).toHaveBeenCalledWith([jasmine.any(String)]);
        } else if (index === 6) {
          expect(service.jobsService.jobPanelItem.open).toHaveBeenCalled();
        }
      });
    });

    it('should set language change actions for mobile navigation zero options', () => {
      // Arrange
      const dropdowns = service.createMobileNavZero();
      // @ts-ignore
      const langsDropdown = dropdowns[1].options[1].options;
      // Act
      service.setActionsForMobileNavZeroOptions(dropdowns);
      // Assert
      langsDropdown.forEach((item, index) => {
        item.click();
        const langCode = enTranslations.global.langAbbreviations[index];
        expect(service.utilService.changeLanguage).toHaveBeenCalledWith(makeLowerCase(langCode));
        expect(service.utilService.currentLang).toEqual(item.text);
      });
    });
  });

  describe('addIdsToMobileNavZeroOptions', () => {
    it('should add IDs to mobile navigation zero options', () => {
      // Arrange
      const dropdowns = service.createMobileNavZero();
      // Act
      service.addIdsToMobileNavZeroOptions(dropdowns);
      // Assert
      const expectedIds = ['logoOption', 'langsOption', 'takeSurveyOption', 'manageVideosOption', 'connectionHubOption', 'progressTrackerOption', 'termsOfServiceOption', 'privacyPolicyOption', 'myAccountOption'];
      dropdowns.forEach((dropdown, index) => {
        expect(dropdown.id).toEqual(ENV.idPrefix.nav + expectedIds[index]);
      });
    });
  });

  describe('createMobileNavOne', () => {
    it('should create mobile navigation one with the correct structure', () => {
      // Arrange & Act
      const navOne = service.createMobileNavOne();
      // Assert
      expect(navOne).toBeInstanceOf(Array);
      navOne.forEach((dropdown) => {
        expect(dropdown).toBeInstanceOf(WMLInfiniteDropdownZeroProps);
      })
      navOne.slice(1).forEach((dropdown) => {
        expect(dropdown).toBeInstanceOf(WMLInfiniteDropdownZeroProps);
      });
    });
  });

  describe('setActionsForMobileNavOneOptions', () => {
    it('should set click actions for mobile navigation one options', () => {
      // Arrange
      const dropdowns = service.createMobileNavOne();
      // Act
      service.setActionsForMobileNavOneOptions(dropdowns);
      // Assert
      dropdowns.slice(2).forEach((dropdown, index) => {
        const option = dropdown.options[0];
        option.click();
        if (index === 4) {
          expect(service.accountsService.signOutViaFirebase).toHaveBeenCalled();
        } else {
          expect(service.utilService.router.navigate).toHaveBeenCalledWith([jasmine.any(String)]);
        }
      });
    });

    it('should set language change actions for mobile navigation one options', () => {
      // Arrange
      const dropdowns = service.createMobileNavOne();
      // @ts-ignore
      const langsDropdown = dropdowns[1].options[1].options;
      // Act
      service.setActionsForMobileNavOneOptions(dropdowns);
      // Assert
      langsDropdown.forEach((item, index) => {
        item.click();
        const langCode = enTranslations.global.langAbbreviations[index];
        expect(service.utilService.changeLanguage).toHaveBeenCalledWith(makeLowerCase(langCode));
        expect(service.utilService.currentLang).toEqual(item.text);
      });
    });
  });

  describe('addIdsToMobileNavOneOptions', () => {
    it('should add IDs to mobile navigation one options', () => {
      // Arrange
      const dropdowns = service.createMobileNavOne();
      // Act
      service.addIdsToMobileNavOneOptions(dropdowns);
      // Assert
      const expectedIds = ['dashboardOption', 'profileOption', 'billingOption', 'dataOption', 'signOutOption'];
      dropdowns.forEach((dropdown, index) => {
        expect(dropdown.id).toEqual(ENV.idPrefix.nav + expectedIds[index]);
      });
    });
  });

});
