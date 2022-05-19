<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use app\Models\Reservation;
use App\Traits\UtilsTrait;
use App\Task;
use App\Models\User;
use App\Http\Requests\StripeRequest;
use App\Repositories\UserRepository;
use Illuminate\Support\Facades\Http;
class ReservationController extends Controller
{
    use UtilsTrait;
    public function __construct(User $user, UserRepository $userRepository) {
        $this->user = $user;
        $this->userRepository = $userRepository;
    }
    public function usergetsessionid(StripeRequest $request){
        $stripe = new \Stripe\StripeClient(env('STRIPE_API_KEY'));
        $total = $request['total_price'];
        $checkout = $stripe->checkout->sessions->create([
            'success_url' => 'http://poliserva.jals.es:4200/#/',
            'cancel_url' => 'http://poliserva.jals.es:4200/#/reservation',
            'line_items' => [
              [
                'price_data' => [
                    'currency' => 'eur',
                    'unit_amount' => $total,
                    'product_data' => [
                        'name'=> "Pista deportiva",
                    ],
                ] ,
                'quantity' => 1,
              ],
            ],
            'mode' => 'payment',
        ]);
        return $checkout;
    }
    public function insertReservation(Request $request){

        $user = $this->userRepository->getUser();
        $request['id_user']=$user['id'];

        $response = HTTP::withHeaders([
            "id_user"=>$request['id_user'],
            "id_court"=>$request->input('id_court'),
            "date"=>$request->input('date'),
            "hini"=>$request->input('hini'),
            "hfin"=>$request->input('hfin'),
            "total_price"=>$request->input('total_price')
        ])->acceptJson()->post("http://172.29.0.12:8003/api/reservations/")->json();

        return $response;

    }
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        //
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        //
    }
}
