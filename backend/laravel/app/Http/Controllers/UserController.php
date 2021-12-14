<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\Http;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use App\Models\User;
use App\Http\Requests\LoginRequest;
use App\Http\Requests\RegisterRequest;
use App\Repositories\UserRepository;
use Symfony\Component\HttpFoundation\Response;
use App\Traits\ApiResponseTrait;
use App\Traits\UtilsTrait;
use JWTAuth;
use Illuminate\Support\Facades\Session;

class UserController extends Controller
{

    protected User $user;
    use ApiResponseTrait;
    use UtilsTrait;

    public function __construct(User $user, UserRepository $userRepository) {
        $this->user = $user;
        $this->userRepository = $userRepository;
    }

    public function login(LoginRequest $request) {

        try{
            if(session('token')){
                $data = self::decode(session('token'));
                $array = json_decode(json_encode($data), True);
                $uuid = $array['uuid'];
            }else{
                // $email = $request->only('email');
                // $response = HTTP::acceptJson()->post("http://localhost:3000/api/users/getRole")->json();

                $response = HTTP::withHeaders(['uuid' => 'asdasd'])->acceptJson()->post("http://localhost:3000/api/users/getrole")->json();

                return $response;
                // print_r($response);
            }
        }catch(Exception $e){

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
            // $dataLogin = $request->only(
            //     'mail',
            //     'pass'
            // );
            
            $user = $this->userRepository->register($dataReq);
            
            if($user){
                $token = self::encode();
               if( $token ){
                    $data =  $this->respondWithToken($token);
                    session(['token'=>$data]);
                    Session::save();
                    // $_SESSION['token'] = $data;
                    // $token1 = $_SESSION['token']; 
                    // $this->login($dataLogin);
                    return redirect()->action([UserController::class,'login'],['mail'=>$dataReq['mail'],'pass'=>$dataReq['pass']]);
                    // return self::apiResponseSuccess(null, 'Successfully registered', Response::HTTP_OK);
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
        $code = $_SESSION['code'];
        $obj = json_encode($arr);
        if($codeVerify != $code){
            return self::apiServerError($e->getMessage());
        }else{
            $this->register($obj);
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
            self::dataMail($arrMail);
        }
    }
    protected function respondWithToken($token)
    {
        $data = [
            'access_token'  => $token,
            'token_type'    => 'bearer',
            'expires_in'    => $this->guard()->factory()->getTTL() * 60 * 24 * 30,
            'user'          => $this->guard()->user()
        ];
        return $data['access_token'];
    }

    public function guard()
    {
        return auth()->guard('api');
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
