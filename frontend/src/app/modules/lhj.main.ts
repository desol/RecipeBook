// Libraries
import 'hammerjs';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { UrlSerializer } from '@angular/router'
import { FlexLayoutModule } from '@angular/flex-layout';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule, Component, ViewContainerRef } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MdButtonModule, MdCardModule, MdMenuModule, MdToolbarModule, MdProgressBarModule,
  MdIconModule, MdInputModule, MdSnackBarModule, MdTooltipModule
} from '@angular/material';

// Components
import { MenuComponent } from '../components/lhj.menu';
import { LoginComponent } from '../components/lhj.login';
import { AccountComponent } from '../components/lhj.account';
import { ProgressBarComponent } from '../components/lhj.progressbar'

// Services
import { Session } from '../services/session';
import { LowerCaseUrlSerializer } from '../services/urlserializer';

// Modules
import { AppRouting } from './lhj.routing';

@Component({
  selector: 'app-lhj-main',
  templateUrl: '../pages/lhj.main.html'
})

export class MainComponent {

};

@NgModule({
  declarations: [
    MainComponent,
    MenuComponent,
    LoginComponent,
    AccountComponent,
    ProgressBarComponent,
  ],
  imports: [
    AppRouting,
    FormsModule,
    MdMenuModule,
    MdIconModule,
    MdCardModule,
    MdInputModule,
    BrowserModule,
    MdButtonModule,
    MdTooltipModule,
    MdToolbarModule,
    MdSnackBarModule,
    HttpClientModule,
    FlexLayoutModule,
    MdProgressBarModule,
    BrowserAnimationsModule,
  ],
  providers: [
    Session,
    {
      provide: UrlSerializer,
      useClass: LowerCaseUrlSerializer
    }
  ],
  bootstrap: [
    MainComponent
  ]
})
export class LHJ { }
