import { Component } from '@angular/core';
// import { trigger, state, style, transition, animate, keyframes } from '@angular/animations';

@Component({
  selector: 'router-outlet',
  templateUrl: '../pages/lhj.login.html',
  styleUrls: ['../styles/lhj.login.css'],
  // animations: [
  //   trigger('TitleAnimation', [
  //     state('small', style({
  //       tranform: 'scale(1)',
  //     })),
  //     state('large', style({
  //       transfrom: ('scale(1.2')
  //     })),
  //     transition('small => large', animate('300ms ease-in'))
  //   ])
  // ]
})
export class Login {
  title = 'Daddy is a big ol stinker.';
}
// function AnimateTitle() {

// }
