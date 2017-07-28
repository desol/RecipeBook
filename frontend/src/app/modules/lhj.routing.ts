import { NgModule } from '@angular/core';
import { RouterModule, Routes, DefaultUrlSerializer } from '@angular/router';

import { LoginComponent } from '../components/lhj.login';
import { AccountComponent } from '../components/lhj.account';

const AppRoutes: Routes = [
  { path: '', component: LoginComponent },
  { path: 'login', component: LoginComponent },
  { path: 'account', component: AccountComponent },
  { path: '**', component: AccountComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(AppRoutes, { useHash: true })],
  exports: [RouterModule]
})

export class AppRouting { }
