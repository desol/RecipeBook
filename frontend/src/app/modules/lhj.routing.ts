import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { Login } from '../components/lhj.login';
import { Account } from '../components/lhj.account';

import { AuthGuard } from '../services/authguard';
import { RootPathGuard } from '../services/rootpathguard';

const AppRoutes: Routes = [
  { path: '', component: Login, canActivate: [RootPathGuard] },
  { path: 'Login', component: Login },
  { path: 'Account', component: Account, canActivate: [AuthGuard] },
];

@NgModule({
  imports: [RouterModule.forRoot(AppRoutes, {useHash: true})],
  exports: [RouterModule]
})

export class AppRouting { }
