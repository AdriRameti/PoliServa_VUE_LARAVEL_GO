<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\Http;
use Illuminate\Http\Request;
use App\Models\User;
use App\Http\Requests\LoginRequest;
use App\Http\Requests\RegisterRequest;
use App\Repositories\UserRepository;
use Symfony\Component\HttpFoundation\Response;
use App\Traits\ApiResponseTrait;
use App\Traits\Google2FATrait;
use App\Traits\UtilsTrait;
use Illuminate\Support\Facades\Session;

class UserController extends Controller
{

    protected User $user;
    use ApiResponseTrait;
    use UtilsTrait;
    use Google2FATrait;

    public function __construct(User $user, UserRepository $userRepository) {
        $this->user = $user;
        $this->userRepository = $userRepository;
    }

    public function enable2fa() {
        return self::en2fa();
    }

    public function disable2fa() {
        return self::dis2fa();
    }

    public function check2fa(Request $request) {

        $digits = $request->only('one_time_password');

        return self::checkOTP($digits['one_time_password']);
    }

    public function login(LoginRequest $request) {

        try{
            if(session('token')){
                $data = self::decode(session('token'));

                $array = json_decode(json_encode($data), True);

                $uuid = $array['uuid'];

                $response = HTTP::withHeaders(['uuid' => $uuid])->acceptJson()->post("http://localhost:3000/api/users/getrole")->json();

                return $response;
            }else{

                $data = $request->validated();

                $uuid = $this->userRepository->login($data);

                if ($uuid == "user not found") {
                    return self::apiServerError('user not found');
                } else if ($uuid == "password don't match") {
                    return self::apiServerError("password don't match");
                } else {
                    $response = HTTP::withHeaders(['uuid' => $uuid])->acceptJson()->post("http://localhost:3000/api/users/getrole")->json();
                }

                return $response;
            }
        }catch(Exception $e){
            return self::apiServerError($e->getMessage());
        }

    }

    public function register(RegisterRequest $request) {
        try{
            $dataReq = $request->only(
                'name',
                'surnames',
                'mail',
                'pass'
            );

            $user = $this->userRepository->register($dataReq);
            if($user){
                $token = self::encode();
               if( $token ){

                    session(['token'=>$token]);

                    Session::save();

                    return redirect()->action([UserController::class,'login'],['mail'=>$dataReq['mail'],'pass'=>$dataReq['pass']]);
                } 
            }
        }catch(\Exception $e){
            return self::apiServerError($e->getMessage());
        }
    }

    public function mailRegister($infouser){
        $arr = array(
            'name' => $infouser[0],
            'surnames' => $infouser[1],
            'mail' => $infouser[2],
            'pass' => $infouser[3],
        );
        $codeVerify = $infouser[4];
        $code = session('code');
        if($codeVerify != $code){
            return self::apiServerError($e->getMessage());
        }else{
            return redirect()->action([ UserController::class, 'register' ],[ 'name' => $arr['name'] , 'surnames' => $arr['surnames'] , 'mail' => $arr['mail'] ,'pass' => $arr['pass']  ]);
        }
    }

    public function sendMailRegister($info){
        $mail = $info[0];
        $type = 'register';
        $arrMail = array();
        array_push($arrMail,$mail);
        array_push($arrMail,$type);
        if (!$mail && !$type){
            return self::apiServerError($e->getMessage());
        }else{
            return self::dataMail($arrMail);
        }
    }

    // protected function respondWithToken($token)
    // {
    //     $data = [
    //         'access_token'  => $token,
    //         'token_type'    => 'bearer',
    //         'expires_in'    => $this->guard()->factory()->getTTL() * 60 * 24 * 30,
    //         'user'          => $this->guard()->user()
    //     ];
    //     return $data['access_token'];
    // }

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
