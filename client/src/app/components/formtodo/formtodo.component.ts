import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-formtodo',
  templateUrl: './formTodo.component.html',
})
export class FormtodoComponent implements OnInit {
  form = new FormGroup({
    title: new FormControl<string>(''),
  });

  constructor() {}

  ngOnInit(): void {}

  submit() {
    console.log(this.form.value);
  }
}
