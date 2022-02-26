import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatTableModule } from '@angular/material/table';
import { MatTabsModule } from '@angular/material/tabs';
import { SharedModule } from '../../../shared/shared.module';
import { AppComponentDetailComponent } from './app-component-detail.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    MatTabsModule,
    MatCardModule,
    MatTableModule,
    SharedModule
  ],
  declarations: [
    AppComponentDetailComponent
  ],
})
export class AppComponentDetailModule { }