import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ITodo } from 'src/app/models/itodo';
import { TodoService } from 'src/app/services/todo.service';

@Component({
  selector: 'app-todolist',
  templateUrl: './todolist.component.html',
})
export class TodolistComponent implements OnInit {
  // todos$: Observable<ITodo[]>;
  todos: ITodo[] = [
    {
      userId: 1,
      id: 1,
      title: 'Wash dishes',
      completed: false,
    },
    {
      userId: 1,
      id: 1,
      title: 'walk the dog',
      completed: true,
    },
  ];
  loading = false;

  constructor(private todosService: TodoService) {}

  ngOnInit(): void {
    // this.loading = true;
    // this.todos$ = this.todosService.getTodos();
  }
}
