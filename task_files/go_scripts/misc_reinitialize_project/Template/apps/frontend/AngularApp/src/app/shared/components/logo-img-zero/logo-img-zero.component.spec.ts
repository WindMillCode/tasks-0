// testing
import { ComponentFixture, discardPeriodicTasks, fakeAsync, flush } from '@angular/core/testing';


import { wmlTestUtils } from '@core/utility/test-utils/test-utils';
import { resetImports } from '@core/utility/test-utils/mock-imports';
import { resetProviders } from '@core/utility/test-utils/mock-providers';
import { resetDeclarations } from '@core/utility/test-utils/mock-declarations';




// rxjs
import { Subject, of } from 'rxjs';

import { LogoImgZeroComponent, LogoImgZeroProps } from './logo-img-zero.component';
import { ENV } from '@env/environment';
import { AccountsServiceUser } from '@shared/services/accounts/accounts.service';

describe('LogoImgZeroComponent', () => {
  let cpnt: LogoImgZeroComponent;
  let fixture: ComponentFixture<LogoImgZeroComponent>;

  beforeEach(async () => {
    resetImports()
    resetProviders()
    resetDeclarations()



    await wmlTestUtils.configureTestingModuleForComponents(LogoImgZeroComponent);


    ({fixture, cpnt} =  wmlTestUtils.grabComponentInstance(LogoImgZeroComponent));
    fixture.detectChanges()

    spyOn(cpnt.cdref,"detectChanges")
  })

  describe("init", () => {
    it("should create", () => {
      expect(cpnt).toBeTruthy();
    });

    it("should have all values initialize properly", () => {
      expect(cpnt.myClass).toEqual('LogoImgZeroView ');
      expect(cpnt.homeLocation).toEqual(ENV.nav.urls.landing);
    });

    it("should have all properties be the correct class instance", () => {
      expect(cpnt.ngUnsub).toBeInstanceOf(Subject<void>);
      expect(cpnt.props).toBeInstanceOf(LogoImgZeroProps);
    });
  });

  describe("navToHome", () => {
    it("navigates to the home location", () => {
      // Arrange
      cpnt.homeLocation = ENV.nav.urls.home;
      spyOn(cpnt.utilService.router, 'navigate');

      // Act
      cpnt.navToHome();

      // Assert
      expect(cpnt.utilService.router.navigate).toHaveBeenCalledWith([ENV.nav.urls.home]);
    });
  });

  describe("listenForUser", () => {

    it(`when called |
    if currentUser is present |
    sets homeLocation to ENV.nav.urls.home`, fakeAsync(() => {
      // Arrange
      const currentUserMock = new AccountsServiceUser({
        user:{
          // @ts-ignore
          id: "user1"
        }
      });
      cpnt.accountsService.currentUser = currentUserMock;
      cpnt.accountsService.listenForUser = jasmine.createSpy().and.returnValue(of(cpnt.accountsService.currentUser));

      let nextFn = () => {
        // Assert
        expect(cpnt.homeLocation).toEqual(ENV.nav.urls.home);
        discardPeriodicTasks();
      };

      // Act
      cpnt.listenForUser().subscribe({next:nextFn});
      flush();
    }));

    it(`when called |
    if currentUser is not present |
    does not change homeLocation from default`, fakeAsync(() => {
      // Arrange
      cpnt.accountsService.currentUser = null; // Ensure no user is logged in
      const defaultHomeLocation = cpnt.homeLocation; // Capture the default home location
      cpnt.accountsService.listenForUser = jasmine.createSpy().and.returnValue(of(cpnt.accountsService.currentUser));


      let nextFn = () => {
        // Assert
        expect(cpnt.homeLocation).toEqual(defaultHomeLocation);
        discardPeriodicTasks();
      };

      // Act
      cpnt.listenForUser().subscribe({next:nextFn});
      flush();
    }));
  });



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
