import { NgModule } from '@angular/core';
import { RouterModule, Routes, DefaultUrlSerializer } from '@angular/router';

import { LoginComponent } from '../components/lhj.login';
import { AccountComponent } from '../components/lhj.account';

import { AuthGuard } from '../services/authguard';
import { RootPathGuard } from '../services/rootpathguard';

const AppRoutes: Routes = [
  { path: '', component: LoginComponent, canActivate: [RootPathGuard] },
  { path: 'login', component: LoginComponent },
  { path: 'account', component: AccountComponent, canActivate: [AuthGuard] },
  { path: '**', component: AccountComponent, canActivate: [AuthGuard] },
];

@NgModule({
  imports: [RouterModule.forRoot(AppRoutes, { useHash: true })],
  exports: [RouterModule]
})

export class AppRouting { }
