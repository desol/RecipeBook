// Libraries
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { UrlSerializer } from '@angular/router'
import { FlexLayoutModule } from '@angular/flex-layout';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule, Component, ViewContainerRef } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

// Pages
import { LoginComponent } from '../components/lhj.login';
import { AccountComponent } from '../components/lhj.account';

// Services
import { AuthGuard } from '../services/authguard';
import { RootPathGuard } from '../services/rootpathguard';
import { LowerCaseUrlSerializer } from '../services/urlserializer';

// Modules
import { AppRouting } from '../modules/lhj.routing';

@Component({
  selector: 'app-lhj-main',
  templateUrl: '../pages/lhj.main.html'
})

export class MainComponent {


  constructor() {
  }

  // ngOnInit() {

  // }
};

@NgModule({
  declarations: [
    MainComponent,
    LoginComponent,
    AccountComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRouting,
    BrowserAnimationsModule
  ],
  providers: [
    AuthGuard,
    RootPathGuard,
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
