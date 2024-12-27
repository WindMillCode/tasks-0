// angular
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

// wml components
import { WMLComponentsModule } from './wml-components/wml-components.module';

// i18n
import {  TranslateModule, TranslatePipe } from '@ngx-translate/core';
import { OverlayLoadingComponent } from './components/overlay-loading/overlay-loading.component';

// misc
import { SampleCpntComponent } from './components/sample-cpnt/sample-cpnt.component';
import { ConfirmDialogZeroComponent } from './components/confirm-dialog-zero/confirm-dialog-zero.component';
import { ReactiveFormsModule } from '@angular/forms';
import { LogoImgZeroComponent } from './components/logo-img-zero/logo-img-zero.component';



let components = [
  ConfirmDialogZeroComponent,
  SampleCpntComponent,
  OverlayLoadingComponent,
  LogoImgZeroComponent
]

let modules = [
  TranslateModule,
  CommonModule,
  WMLComponentsModule,
  ReactiveFormsModule,
  RouterModule,
]
@NgModule({
  imports:[
    ...modules,
  ],
  exports: [
    ...components,
    ...modules
  ],
  providers:[
    {provide:TranslatePipe,useExisting:TranslatePipe}
  ],
  declarations: [
    ...components
  ]
})
export class SharedModule { }
