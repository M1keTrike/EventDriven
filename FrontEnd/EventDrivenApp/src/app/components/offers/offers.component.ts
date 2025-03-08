import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { WebSocketSubject } from 'rxjs/webSocket';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
import { environment } from '../../../environments/enviroment';

interface Offer {
  id: number;
  user_id: number;
  Product: string;
  Price: number;
  Time: string;
}

@Component({
  selector: 'app-offers',
  templateUrl: './offers.component.html',
  styleUrls: ['./offers.component.css'],
})
export class OffersComponent implements OnInit {
  offers: Offer[] = [];
  offerForm: FormGroup;
  userId: number;
  private socket$: WebSocketSubject<any>;

  constructor(
    private http: HttpClient,
    private fb: FormBuilder,
    private toastr: ToastrService
  ) {
    this.userId =
      Number(sessionStorage.getItem('userId')) ||
      Math.floor(Math.random() * 1000);
    sessionStorage.setItem('userId', this.userId.toString());

    this.offerForm = this.fb.group({
      product: [''],
      price: [''],
      time: [''],
      user_id: [this.userId],
    });

    this.socket$ = new WebSocketSubject(
      `${environment.wsUrl}/ws?destinationID=${this.userId}`
    );
  }

  ngOnInit(): void {
    this.fetchOffers();
    this.socket$.subscribe(
      (message) => {
        if (message) {
          this.toastr.success(
            'Nueva oferta creada, actualizando lista...',
            'Notificación'
          );
          this.fetchOffers();
        }
      },
      (err) => console.error('WebSocket Error:', err)
    );
  }

  fetchOffers(): void {
    this.http
      .get<{ offers: Offer[] }>(`${environment.apiUrl}/offers`)
      .subscribe((data) => {
        this.offers = data.offers;
      });
  }

  createOffer(): void {
    const newOffer = { ...this.offerForm.value, user_id: this.userId };
    this.http.post(`${environment.apiUrl}/offers`, newOffer).subscribe(() => {
      this.offerForm.reset({ user_id: this.userId });
    });
  }
}
