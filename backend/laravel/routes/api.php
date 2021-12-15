<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\CourtsController;
use App\Http\Controllers\ReservationController;
use App\Http\Controllers\SportsController;
use App\Http\Controllers\UserController;
/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});
Route::get('users/login', [UserController::class, "login"]);
Route::post('users/login', [UserController::class, "login"]);
Route::get('users/register', [UserController::class, "register"]);
Route::post('users/register', [UserController::class, "register"]);
Route::post('users/mailRegister',[UserController::class,"mailRegister"]);
Route::post('users/sendMailRegister',[UserController::class,"sendMailRegister"]);
Route::resource('users', UserController::class);
Route::resource('reservation',ReservationController::class);
Route::resource('courts',CourtsController::class);
Route::resource('sports',SportsController::class);