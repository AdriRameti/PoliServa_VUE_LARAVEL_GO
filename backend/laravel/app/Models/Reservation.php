<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Reservation extends Model
{
    use HasFactory;

    protected $fillable = ['date', 'hini','hfin', 'total_price'];

    protected $hidden = ['created_at', 'updated_at'];

    public function id_user() {
        return $this->belongsTo(User::class);
    }

    public function id_court() {
        return $this->belongsTo(Court::class);
    }
}
