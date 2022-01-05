<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\Http;
use Illuminate\Http\Request;
use App\Models\User;
use App\Http\Resources\UserResource;
use App\Http\Requests\LoginRequest;
use App\Http\Requests\RegisterRequest;
use App\Http\Requests\UpdateRequest;
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

    public function login(LoginRequest $request): UserResource|Response {

        try{
            if(session('token')){
                $uuid = self::getUuid();
                $user = $this->userRepository->getUser();
                $fullName = $user->fullName;
                $response = HTTP::withHeaders(['uuid' => $uuid])->acceptJson()->post("http://localhost:3000/api/users/getrole")->json();
                $response+=array('fullName'=>$fullName);
                $response+=array('token'=>session('token'));
                return $this->userResponse($response);
            }else{

                $data = $request->validated();

                $uuid = $this->userRepository->login($data);

                if ($uuid == "user not found") {
                    return self::apiResponseLogin('user not found');
                } else if ($uuid == "password don't match") {
                    return self::apiResponseLogin("password don't match");
                } else {
                    $response = HTTP::withHeaders(['uuid' => $uuid])->acceptJson()->post("http://localhost:3000/api/users/getrole")->json();
                }

                session(['uuid'=>$uuid]);
                Session::save();

                if (session('uuid')) {
                    $token = self::encode();

                    if( $token ) {
                        session(['token'=>$token]);
                        Session::save();
                    }

                    $user = $this->userRepository->getUser();
                    $fullName = $user->fullName;

                    $response+=array('fullName'=>$fullName);
                    $response+=array('token'=>$token);
                }
                
                return $this->userResponse($response);
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
    public function update(UpdateRequest $request)
    {
        if($this->userRepository->getUser()){
            $user = $this->userRepository->getUser();
            $user->name = $request->name;
            $user->surnames = $request->surnames;
            $user->mail = $request->mail;
            $user->pass = $request->pass;
            $user->img = $request->img;
            $user->save();
            print_r($user);
            // return $this->userRepository->getUser();
        }else{
            self::apiResponseError($e->getMessage());
        }
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function deleteUser(Request $request){
        if($this->userRepository->check2fa()){
            $result = self::checkOTP($request->input('one_time_password'));
            if($result === 'verified'){
                $user = $this->userRepository->getUser();
                $user->delete();
                return self::apiResponseSuccess(null);
            }
        }else{
             $user = $this->userRepository->getUser();
             $type = 'delete';
             $arrMail = array();
             array_push($arrMail,$user['mail']);
             array_push($arrMail,$type);
             if (!$user['mail'] && !$type){
                 return self::apiServerError($e->getMessage());
             }else{
                 return self::dataMail($arrMail);
             }
        }
    }
    public function destroy(Request $request)
    {
        $codeVerify = $request->input('codeVeri');
        $code = session('code');
        if($codeVerify != $code){
            return self::apiServerError($e->getMessage());
        }else{
            $user = $this->userRepository->getUser();
            $user->delete();
            return self::apiResponseSuccess(null);
        }
    }

    private function userResponse($response): UserResource {
        // $user = $this->userRepository->getUser();
        // $fullName = $user->fullName;
        
        return new UserResource($response);
    }
}
