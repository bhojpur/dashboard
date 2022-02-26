import { Component, OnInit } from '@angular/core';
import { AppComponent } from 'src/app/types/types';
import { ComponentsService } from 'src/app/components/components.service';
import { ActivatedRoute } from '@angular/router';
import * as yaml from 'js-yaml';
import { ThemeService } from 'src/app/theme/theme.service';

@Component({
  selector: 'app-app-component-detail',
  templateUrl: './app-component-detail.component.html',
  styleUrls: ['./app-component-detail.component.scss']
})
export class AppComponentDetailComponent implements OnInit {

  private name: string | undefined;
  public component: any;
  public componentManifest!: string;
  public loadedComponent = false;

  constructor(
    private route: ActivatedRoute,
    private componentsService: ComponentsService,
    private themeService: ThemeService,
  ) { }

  ngOnInit(): void {
    this.name = this.route.snapshot.params.name;
    if (typeof this.name !== 'undefined') {
      this.getComponent(this.name);
    }
  }

  getComponent(name: string): void {
    this.componentsService.getComponent(name).subscribe((data: AppComponent) => {
      this.component = data;
      this.componentManifest = (typeof data.manifest === 'string') ?
        data.manifest : yaml.dump(data.manifest);
      this.loadedComponent = true;
    });
  }

  isDarkTheme() {
    return this.themeService.isDarkTheme();
  }
}