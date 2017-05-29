import { Component } from '@angular/core';
import { trigger, state, style, transition, animate, keyframes } from '@angular/animations';

@Component({
  selector: 'app-root',
  templateUrl: '../Pages/lhj.login.html',
  styleUrls: ['../Styles/lhj.login.css'],
  animations: [
    trigger('TitleAnimation', [
      state('small', style({
        tranform: 'scale(1)',
      })),
      state('large', style({
        transfrom: ('scale(1.2')
      })),
      transition('small => large', animate('300ms ease-in'))
    ])
  ]
})
export class Login {
  title = 'Kevin is a big ol stinker.';
}
function AnimateTitle() {

}
