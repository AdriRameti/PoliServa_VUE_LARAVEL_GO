<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use app\Models\Reservation;
use App\Traits\UtilsTrait;
use App\Task;
use App\Models\User;
use App\Repositories\UserRepository;
use Illuminate\Support\Facades\Http;
class ReservationController extends Controller
{
    use UtilsTrait;
    public function __construct(User $user, UserRepository $userRepository) {
        $this->user = $user;
        $this->userRepository = $userRepository;
    }
    public function insertReservation(Request $request){
            $user = $this->userRepository->getUser();
            $request['id_user']=$user['id'];
            $response = HTTP::withHeaders([
                "id_user"=>$request['id_user'],
                "id_court"=>$request['id_court'],
                "date"=>$request['date'],
                "hini"=>$request['hini'],
                "hfin"=>$request['hfin'],
                "total_price"=>$request['total_price']
            ])->acceptJson()->post("http://localhost:3000/api/reservations/")->json();
        return $response;
    //    $uuid = self::getUuid();


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
