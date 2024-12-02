// testing
import { ComponentFixture } from '@angular/core/testing';


import { wmlTestUtils } from '@core/utility/test-utils/test-utils';
import { resetImports } from '@core/utility/test-utils/mock-imports';
import { resetProviders } from '@core/utility/test-utils/mock-providers';
import { resetDeclarations } from '@core/utility/test-utils/mock-declarations';




// rxjs
import { Subject } from 'rxjs';

import { LegalDocPageComponent } from './doc-zero-page.component';
import { ENV } from '@env/environment';


describe('LegalDocPageComponent', () => {
  let cpnt: LegalDocPageComponent;
  let fixture: ComponentFixture<LegalDocPageComponent>;

  beforeEach(async () => {
    resetImports()
    resetProviders()
    resetDeclarations()
    wmlTestUtils.mock.mockUtilService.router.url =ENV.nav.urls.privacyPolicyPage
    await wmlTestUtils.configureTestingModuleForComponents(LegalDocPageComponent);
    ({fixture, cpnt} =  wmlTestUtils.grabComponentInstance(LegalDocPageComponent));

    fixture.detectChanges()
    spyOn(cpnt.cdref,"detectChanges")
  })

  describe("init", () => {

    it("should create", () => {
      expect(cpnt).toBeTruthy()
    })

    it("should have all values initalize properly", () => {
      expect(cpnt.myClass).toEqual('LegalDocPageView ')
    })

    it("should have all properties be the correct class instance", () => {
      expect(cpnt.ngUnsub).toBeInstanceOf(Subject<void>)
    })
  })



  describe("ngOnDestroy",()=>{

    beforeEach(()=>{
      spyOn(cpnt.ngUnsub,'next')
      spyOn(cpnt.ngUnsub,'complete')
    })

    it(` when called |
     as appropriate |
     does the required action `,()=>{
        // act
        cpnt.ngOnDestroy();

        // assert
        expect(cpnt.ngUnsub.next).toHaveBeenCalled();
        expect(cpnt.ngUnsub.complete).toHaveBeenCalled();
    })
  })
});
