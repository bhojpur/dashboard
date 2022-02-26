import { Component, OnInit, OnDestroy } from '@angular/core';
import { ComponentsService } from 'src/app/components/components.service';
import { ScopesService } from 'src/app/scopes/scopes.service';
import { AppComponent } from 'src/app/types/types';

@Component({
  selector: 'app-components',
  templateUrl: 'app-components.component.html',
  styleUrls: ['app-components.component.scss'],
})
export class AppComponentsComponent implements OnInit, OnDestroy {

  public components: AppComponent[] = [];
  public componentsLoaded = false;
  public displayedColumns: string[] = ['img', 'name', 'status', 'age', 'created'];
  private intervalHandler: any;

  constructor(
    private componentsService: ComponentsService,
    private scopesService: ScopesService
  ) { }

  ngOnInit(): void {
    this.getComponents();

    this.intervalHandler = setInterval(() => {
      this.getComponents();
    }, 10000);

    this.scopesService.scopeChanged.subscribe(() => {
      this.getComponents();
    });
  }

  ngOnDestroy(): void {
    clearInterval(this.intervalHandler);
  }

  getComponents(): void {
    this.componentsService.getComponents().subscribe((data: AppComponent[]) => {
      this.components = data;
      this.components.forEach(component => {
        component.img = this.getIconPath(component.type);
      });
      this.componentsLoaded = true;
    });
  }

  getIconPath(type: string): string {
    if (type.includes('bindings')) {
      return 'assets/images/bindings.png';
    } else if (type.includes('secretstores')) {
      return 'assets/images/secretstores.png';
    } else if (type.includes('state')) {
      return 'assets/images/statestores.png';
    } else if (type.includes('pubsub')) {
      return 'assets/images/pubsub.png';
    } else if (type.includes('exporters')) {
      return 'assets/images/tracing.png';
    } else {
      return 'assets/images/secretstores.png';
    }
  }
}