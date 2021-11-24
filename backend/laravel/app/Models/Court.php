<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Court extends Model
{
    use HasFactory;

    protected $fillable = ['sector', 'price/h'];
    protected $hidden = ['created_at', 'updated_at'];

    public function id_sport() {
        return $this->belongsTo(Sport::class);
    }

}
