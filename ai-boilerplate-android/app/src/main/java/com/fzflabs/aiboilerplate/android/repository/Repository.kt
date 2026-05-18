package com.fzflabs.aiboilerplate.android.repository

import com.fzflabs.aiboilerplate.android.model.ProductDetailsDto
import com.fzflabs.aiboilerplate.android.util.ApiState

interface Repository {
    suspend fun getProductDetails(): ApiState<ProductDetailsDto>
}