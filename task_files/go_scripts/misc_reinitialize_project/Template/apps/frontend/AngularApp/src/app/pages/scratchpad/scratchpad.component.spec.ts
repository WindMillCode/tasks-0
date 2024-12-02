// testing
import { ComponentFixture } from '@angular/core/testing';
import { wmlTestUtils } from '@core/utility/test-utils/test-utils';
import { resetImports } from '@core/utility/test-utils/mock-imports';
import { resetProviders } from '@core/utility/test-utils/mock-providers';
import { resetDeclarations } from '@core/utility/test-utils/mock-declarations';

// rxjs
import { Subject } from 'rxjs';

import { ScratchpadComponent } from './scratchpad.component';
import { WMLButtonZeroProps } from '@windmillcode/angular-wml-button';


describe('ScratchpadComponent', () => {
  let cpnt: ScratchpadComponent;
  let fixture: ComponentFixture<ScratchpadComponent>;

  beforeEach(async () => {

    resetImports()
    resetProviders()
    resetDeclarations()
    await wmlTestUtils.configureTestingModuleForStandaloneComponents(ScratchpadComponent);




    ({fixture, cpnt} =  wmlTestUtils.grabComponentInstance(ScratchpadComponent));
    fixture.detectChanges()
  })

  describe("init", () => {

    it("should create", () => {
      expect(cpnt).toBeTruthy()
    })

    it("should have all values initalize properly", () => {
      expect(cpnt.myClass).toEqual('ScratchpadView ')
    })

    it("should have all properties be the correct class instance", () => {
      expect(cpnt.ngUnsub).toBeInstanceOf(Subject<void>);
      expect(cpnt.btn).toBeInstanceOf(WMLButtonZeroProps);
      expect(cpnt.btn1).toBeInstanceOf(WMLButtonZeroProps);
      expect(cpnt.btn2).toBeInstanceOf(WMLButtonZeroProps);
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
        expect(cpnt.ngUnsub.next).toHaveBeenCalledWith();
        expect(cpnt.ngUnsub.complete).toHaveBeenCalledWith();
    })
  })
});
