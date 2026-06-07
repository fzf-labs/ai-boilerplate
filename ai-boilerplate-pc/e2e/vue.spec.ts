import { test, expect } from '@playwright/test';

test('renders the AI Boilerplate landing page', async ({ page }) => {
  await page.goto('/');
  await expect(page.locator('h1')).toHaveText('一套系统打通会员、商城与内容运营。');
})
