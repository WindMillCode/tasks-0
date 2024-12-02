// testing
import { wmlTestUtils } from '@core/utility/test-utils/test-utils';
import { resetImports } from '@core/utility/test-utils/mock-imports';
import { resetProviders } from '@core/utility/test-utils/mock-providers';
import { resetDeclarations } from '@core/utility/test-utils/mock-declarations';
import { TestBed } from '@angular/core/testing';

// services
import { UtilityService } from '@core/utility/utility.service';

import { SpecificService } from './specific.service';
import { WMLImage } from '@windmillcode/wml-components-base';
import { ENV } from '@env/environment';
import { GenerateFieldProps } from '@core/utility/form-utils';
import { WMLOptionsZeroProps, WMLOptionZeroItemProps } from '@windmillcode/angular-wml-options';
import { FormArray, FormGroup } from '@angular/forms';
import { CardNumberInputZeroProps, CardNumberInputZeroComponent } from '@shared/components/card-number-input-zero/card-number-input-zero.component';
import { VerifiedZeroProps, VerifiedZeroComponent } from '@shared/components/verified-zero/verified-zero.component';

let formUtils = require('@core/utility/form-utils')

describe('SpecificService', () => {
  let service: SpecificService;
  let utilService:UtilityService

  beforeEach(() => {
    resetImports()
    resetProviders();
    resetDeclarations()
    service = wmlTestUtils.configureTestingModuleForServices(SpecificService)
    utilService =TestBed.inject(UtilityService)
    spyOn(service,"createField").and.callThrough()

    spyOn(formUtils,'createWMLFormZeroPropsField').and.callThrough()


    spyOn(service.utilService.router, 'navigateByUrl')
    spyOn(service.baseService,"createOptionsFormField").and.callThrough()
  });

  describe("init", () => {

    it("should create", () => {
      expect(service).toBeTruthy()
    })

    it("should have all values initalize properly", () => {

    })

    it("should have all properties be the correct class instance", () => {

    })
  })

  describe("logoImg",()=>{
    it(` when called |
    under normal conditions |
    does the required action `,()=>{
      // arrange
      let result
      // act
      result = service.getLogoImg()
      // assert
      expect(result).toBeInstanceOf(WMLImage)
    })

    describe("click",()=>{
      it(` when called |
      under normal conditions |
      does the required action `,()=>{
        // arrange
        let item = service.getLogoImg()
        // act
        item.click()
        // assert
        expect(service.utilService.router.navigateByUrl).
        toHaveBeenCalledWith(ENV.nav.urls.home);

      })
    })
  })

  describe("_manageWebStorage",()=>{
    it(` when called |
    under normal conditions |
    does the required action `,()=>{
      // arrange
      let webStorage = sessionStorage
      let predicate = {
        predicate:()=>{}
      }
      let thisWebStorage = {}
      spyOn(predicate,"predicate")
      spyOn(webStorage,"getItem").and.returnValue(JSON.stringify({a:"a"}))
      spyOn(webStorage,"setItem")
      // act
      service._manageWebStorage(webStorage,thisWebStorage,predicate.predicate)
      // assert
      expect(webStorage.getItem).toHaveBeenCalledWith(ENV.classPrefix.app)
      expect(predicate.predicate).toHaveBeenCalledWith()


    })
  })


  describe('createColorZeroFormField', () => {
    it(` when called |
      under normal conditions |
      does the required action `, () => {
      // arrange
      const props = new GenerateFieldProps<WMLOptionsZeroProps>({
        // @ts-ignore
        colors: [{ color: 'red', value: '1' }],
        fieldParentForm: new FormGroup({
          colorFieldArray:new FormArray([])
        }),
        fieldFormControlName: 'colorFieldArray',
      });
      // act
      // @ts-ignore
      const result = service.createColorZeroFormField(props);

      // assert
      expect(result).toBeInstanceOf(WMLField)
      expect(result.field.custom.props).toBeInstanceOf(WMLOptionsZeroProps);
      expect(result.field.custom.props.options[0]).toBeInstanceOf(WMLOptionZeroItemProps);
      expect(service.baseService.createOptionsFormField).toHaveBeenCalledWith(props);
    });
  });



  describe('createCardNumberInputField', () => {
    it(`when called |
      under normal conditions |
      does the required action`, () => {
      // arrange
      const props = new GenerateFieldProps<CardNumberInputZeroProps>();
      // act
      service.createCardNumberInputField(props);
      // assert
      expect(service.createField).toHaveBeenCalledWith(props, CardNumberInputZeroComponent);
    });
  });



  describe('createVerifiedField', () => {
    it(`when called |
      under normal conditions |
      does the required action`, () => {
      // arrange
      const props = new GenerateFieldProps<VerifiedZeroProps>();
      // act
      service.createVerifiedField(props);
      // assert
      expect(service.createField).toHaveBeenCalledWith(props, VerifiedZeroComponent);
    });
  });

  describe('createField', () => {
    it(`when called |
      under normal conditions |
      does the required action`, () => {
      // arrange
      const props = new GenerateFieldProps<any>();
      const serviceInstance: any= {}
      // act
      let result = service.createField(props, serviceInstance);
      // assert
      expect(formUtils.createWMLFormZeroPropsField).toHaveBeenCalledWith(props)
      expect(result.field.custom).toEqual(jasmine.objectContaining({
        cpnt:serviceInstance,
        props:props.fieldCustomProps
      }))

    });
  });
});
