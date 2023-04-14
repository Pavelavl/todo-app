import { Component, OnInit } from '@angular/core';
import { ITodo } from './models/itodo';
import { TodoService } from './services/todo.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {
  title = 'todo-app';
  todos: ITodo[] = [];

  constructor(private todosService: TodoService) {}

  ngOnInit(): void {
    this.todosService.getTodos().subscribe(todos => {
      this.todos = todos;
    })
  }
}
