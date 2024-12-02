// angular
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, HostBinding   } from '@angular/core';
// services

import { UtilityService } from '@app/core/utility/utility.service';
import { BaseService } from '@core/base/base.service';

// rxjs
import { Subject, defer, delay, finalize, timer } from 'rxjs';

// misc
import { SharedModule } from '@shared/shared.module';
import { generateClassPrefix } from '@windmillcode/wml-components-base';
import { HttpClient } from '@angular/common/http';
import { NavService } from '@shared/services/nav/nav.service';
import { TextEditor } from '@app/core/utility/string-utils';

@Component({
    imports: [
        SharedModule
    ],
    selector: 'scratchpad',
    templateUrl: './scratchpad.component.html',
    styleUrls: ['./scratchpad.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ScratchpadComponent  {

  constructor(
    public cdref:ChangeDetectorRef,
    public http:HttpClient,
    public utilService:UtilityService,
    public baseService:BaseService,
    public navService:NavService,
  ) { }

  classPrefix = generateClassPrefix('Scratchpad')
  idPrefix = generateClassPrefix('scratchpad')


  @HostBinding('class') myClass: string = this.classPrefix(`View`);
  ngUnsub= new Subject<void>()


  runningCalls = 0
  asyncFnThatEndsAtGivenTimeInAnotherPartOfTheCode = (delaySeconds,runSeconds)=>{
    return defer(()=>{
      this.runningCalls +=1
      this.baseService.openOverlayLoading()
      return timer(runSeconds)
      .pipe(
        delay(delaySeconds),
        finalize(()=>{
          this.runningCalls -=1
          if(this.runningCalls === 0){
            this.baseService.closeOverlayLoading()
          }
        })
      )
    })

  }

  ngOnInit(): void {
    debugger
    let a = new TextEditor({
      text:"a b c"
    })
  }

  ngOnDestroy(){
    this.ngUnsub.next();
    this.ngUnsub.complete()
  }

}






