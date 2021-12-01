<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use app\Http\Controllers\CourtsController;
use app\Http\Controllers\ReservationController;
use app\Http\Controllers\SportsController;
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
Route::resouce('reservation',ReservationController::class);
Route::resouce('courts',CourtsController::class);
Route::resouce('sports',SportsController::class);