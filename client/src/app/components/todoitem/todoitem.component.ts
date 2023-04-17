import { Component, Input } from '@angular/core';
import { ITodo } from 'src/app/models/itodo';

@Component({
  selector: 'app-todoitem',
  templateUrl: './todoitem.component.html',
})
export class TodoitemComponent {
  @Input() todoItem: ITodo;
}
