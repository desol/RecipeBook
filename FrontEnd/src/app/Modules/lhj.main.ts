import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { FlexLayoutModule } from '@angular/flex-layout';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule, Component, ViewContainerRef } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

// Pages
import { Login } from '../components/lhj.login';
import { Account } from '../components/lhj.account';

// Services
import { AuthGuard } from '../services/authguard';
import { RootPathGuard } from '../services/rootpathguard';

// Modules
import { AppRouting } from '../modules/lhj.routing';

@Component({
  selector: 'lhj',
  templateUrl: '../pages/lhj.main.html'
})

export class Main {


  constructor() {
  }

  ngOnInit() {

  }
};

@NgModule({
  declarations: [
    Main,
    Login,
    Account,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    BrowserAnimationsModule
  ],
  providers: [
    AppRouting,
    AuthGuard,
    RootPathGuard
  ],
  bootstrap: [
    Main
  ]
})
export class LHJ { }
