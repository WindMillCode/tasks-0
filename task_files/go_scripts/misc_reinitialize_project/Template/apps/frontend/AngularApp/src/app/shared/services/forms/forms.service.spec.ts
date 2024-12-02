// testing
import { wmlTestUtils } from '@core/utility/test-utils/test-utils';
import { resetImports } from '@core/utility/test-utils/mock-imports';
import { resetProviders } from '@core/utility/test-utils/mock-providers';
import { resetDeclarations } from '@core/utility/test-utils/mock-declarations';
import { TestBed } from '@angular/core/testing';

// services
import { UtilityService } from '@core/utility/utility.service';

import { FormsService } from './forms.service';

describe('FormsService', () => {
  let service: FormsService;
  let utilService:UtilityService

  beforeEach(() => {
        resetImports()
    resetProviders();
    resetDeclarations()
    service = wmlTestUtils.configureTestingModuleForServices(FormsService)
    utilService =TestBed.inject(UtilityService)
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
});
