import { Component, Input } from '@angular/core';
import { Process } from '../../models/process';
import { ProcessesService } from '../../services/processes.service';

@Component({
  selector: 'app-conn-card',
  templateUrl: './conn-card.component.html',
  styleUrl: './conn-card.component.css'
})
export class ConnCardComponent {

  @Input() process!: Process;

  constructor(private service: ProcessesService) { }

  onToggleMonitoring(): void {
    this.service.toggleMonitoring(this.process);
  }

  removeProccesFromList() {
    this.service.deleteProcess(this.process.name);
  }
}
