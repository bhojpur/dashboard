import { NgModule } from '@angular/core';
import { PagesComponent } from './pages.component';
import { AboutDialogComponent } from './pages.component';
import { DashboardModule } from './dashboard/dashboard.module';
import { PagesRoutingModule } from './pages-routing.module';
import { CommonModule } from '@angular/common';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { DetailModule } from './dashboard/detail/detail.module';
import { AppComponentsModule } from './app-components/app-components.module';
import { MatListModule } from '@angular/material/list';
import { ConfigurationModule } from './configuration/configuration.module';
import { ControlPlaneModule } from './controlplane/controlplane.module';
import { AppComponentDetailModule } from './app-components/app-component-detail/app-component-detail.module';
import { ConfigurationDetailModule } from './configuration/configuration-detail/configuration-detail.module';
import { OverlayModule } from '@angular/cdk/overlay';
import { MatSelectModule } from '@angular/material/select';
import { FormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';

@NgModule({
  imports: [
    CommonModule,
    PagesRoutingModule,
    DashboardModule,
    AppComponentsModule,
    DetailModule,
    MatSidenavModule,
    MatToolbarModule,
    MatIconModule,
    MatListModule,
    MatButtonModule,
    ConfigurationModule,
    ControlPlaneModule,
    AppComponentDetailModule,
    ConfigurationDetailModule,
    OverlayModule,
    MatSelectModule,
    FormsModule,
    MatDialogModule
  ],
  declarations: [
    PagesComponent,
    AboutDialogComponent
  ],
  entryComponents: [
    AboutDialogComponent
  ]
})
export class PagesModule { }