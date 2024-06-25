import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConnCardComponent } from './conn-card.component';

describe('ConnCardComponent', () => {
  let component: ConnCardComponent;
  let fixture: ComponentFixture<ConnCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ConnCardComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ConnCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
