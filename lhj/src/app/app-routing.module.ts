import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { EventCalendarComponent } from './event-calendar/event-calendar.component';
import { ResumeComponent } from './resume/resume.component';
import { AboutMeComponent } from './about-me/about-me.component';
import { ProfileComponent } from './profile/profile.component';
import { VisitMeComponent } from './visit-me/visit-me.component';

const routes: Routes = [
  { path: '', component: AboutMeComponent },
  { path: 'events', component: EventCalendarComponent },
  { path: 'resume', component: ResumeComponent },
  { path: 'profile', component: ProfileComponent },
  { path: 'visit', component: VisitMeComponent },
  { path: '**', redirectTo: '' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
