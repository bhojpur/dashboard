import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AppComponent } from '../types/types';
import { ScopesService } from '../scopes/scopes.service';

@Injectable({
  providedIn: 'root',
})
export class ComponentsService {

  constructor(
    private http: HttpClient,
    private scopesService: ScopesService
  ) { }

  getComponents(): Observable<AppComponent[]> {
    const scope = this.scopesService.getScope();
    return this.http.get<AppComponent[]>(`/api/components/${scope}`);
  }

  getComponent(name: string): Observable<AppComponent> {
    const scope = this.scopesService.getScope();
    return this.http.get<AppComponent>(`/api/components/${scope}/${name}`);
  }
}