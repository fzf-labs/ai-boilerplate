package com.fzflabs.aiboilerplate.android.network

import com.fzflabs.aiboilerplate.android.model.ProductDetailsDto
import retrofit2.http.GET

interface ApiService {

    @GET("products/1")
    suspend fun getProductDetails(): ProductDetailsDto
}
