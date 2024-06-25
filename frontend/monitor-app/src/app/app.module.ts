import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms'

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { LogMessageComponent } from './components/log-message/log-message.component';
import { LogViewComponent } from './components/log-view/log-view.component';
import { ConnCardComponent } from './components/conn-card/conn-card.component';
import { OptionsViewComponent } from './components/options-view/options-view.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import {LogService } from './services/log.service'
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';


@NgModule({
  declarations: [
    AppComponent,
    LogMessageComponent,
    LogViewComponent,
    ConnCardComponent,
    OptionsViewComponent,
    DashboardComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MatIconModule,
    MatButtonModule,
    FormsModule
  ],
  providers: [
    LogService,
    provideAnimationsAsync()
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
