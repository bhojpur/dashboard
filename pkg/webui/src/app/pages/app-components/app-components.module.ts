import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { AppComponentsComponent } from './app-components.component';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    MatTableModule,
  ],
  declarations: [
    AppComponentsComponent
  ],
})
export class AppComponentsModule { }