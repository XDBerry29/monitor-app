import { LogService } from './../../services/log.service';
import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { Process } from '../../models/process';
import { ProcessesService } from '../../services/processes.service';

@Component({
  selector: 'app-options-view',
  templateUrl: './options-view.component.html',
  styleUrl: './options-view.component.css'
})
export class OptionsViewComponent {
  processes$: Observable<Process[]>;

  constructor(private processService: ProcessesService) {
    this.processes$ = this.processService.processes$;
  }

}
