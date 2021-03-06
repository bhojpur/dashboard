import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AppConfiguration, Instance } from '../types/types';
import { Observable } from 'rxjs';
import { ScopesService } from '../scopes/scopes.service';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ConfigurationsService {

  constructor(
    private http: HttpClient,
    private scopesService: ScopesService,
  ) { }

  getConfigurations(): Observable<AppConfiguration[]> {
    const scope = this.scopesService.getScope();
    return this.http.get<AppConfiguration[]>(`/api/configurations/${scope}`);
  }

  getConfiguration(name: string): Observable<AppConfiguration> {
    const scope = this.scopesService.getScope();
    return this.http.get<AppConfiguration>(`/api/configurations/${scope}/${name}`);
  }

  getConfigurationApps(name: string): Observable<Instance[]> {
    const scope = this.scopesService.getScope();
    return this.http.get<Instance[]>(`/api/instances/${scope}`).pipe(
      map(instances => {
        return instances.filter(instance => instance.config === name);
      })
    );
  }
}