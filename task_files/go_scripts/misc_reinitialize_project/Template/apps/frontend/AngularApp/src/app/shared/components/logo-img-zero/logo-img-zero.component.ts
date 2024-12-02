// angular
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, HostBinding, Input } from '@angular/core';

// services
import { UtilityService } from '@app/core/utility/utility.service';
import { BaseService } from '@core/base/base.service';

// rxjs
import { Subject } from 'rxjs';

// wml-components
import { WMLConstructorDecorator, WMLImage, generateClassPrefix, generateIdPrefix } from '@windmillcode/wml-components-base';

// misc
import { ENV } from '@env/environment';
import { SpecificService } from '@core/specific/specific.service';
import { NavService } from '@shared/services/nav/nav.service';

@Component({
    selector: 'logo-img-zero',
    templateUrl: './logo-img-zero.component.html',
    styleUrls: ['./logo-img-zero.component.scss'],
    changeDetection: ChangeDetectionStrategy.OnPush,
    standalone: false
})
export class LogoImgZeroComponent {
  constructor(
    public cdref:ChangeDetectorRef,
    public specificService:SpecificService,
    public utilService:UtilityService,
    public baseService:BaseService,
    public navService:NavService,
  ) { }

  classPrefix = generateClassPrefix('LogoImgZero');
  idPrefix = generateIdPrefix(ENV.idPrefix.logoImgZero);

  @Input('props') props: LogoImgZeroProps = new LogoImgZeroProps();

  @HostBinding('class') myClass: string = this.classPrefix('View');
  @HostBinding('id') myId: string = this.idPrefix();

  ngUnsub = new Subject<void>();
  homeLocation = ENV.nav.urls.home
  navToHome =(event: Event)=>{
    event.preventDefault();
    this.utilService.router.navigateByUrl(this.homeLocation)
  }



  ngOnInit(): void {
  }

  ngOnDestroy() {
    this.ngUnsub.next();
    this.ngUnsub.complete();
  }
}

@WMLConstructorDecorator
export class LogoImgZeroProps {
  constructor(props: Partial<LogoImgZeroProps> = {}) {}
  backgroundImg:WMLImage
}
